# bash completion for fuzzit                               -*- shell-script -*-

__fuzzit_debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Homebrew on Macs have version 1.3 of bash-completion which doesn't include
# _init_completion. This is a very minimal version of that function.
__fuzzit_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

__fuzzit_index_of_word()
{
    local w word=$1
    shift
    index=0
    for w in "$@"; do
        [[ $w = "$word" ]] && return
        index=$((index+1))
    done
    index=-1
}

__fuzzit_contains_word()
{
    local w word=$1; shift
    for w in "$@"; do
        [[ $w = "$word" ]] && return
    done
    return 1
}

__fuzzit_handle_reply()
{
    __fuzzit_debug "${FUNCNAME[0]}"
    case $cur in
        -*)
            if [[ $(type -t compopt) = "builtin" ]]; then
                compopt -o nospace
            fi
            local allflags
            if [ ${#must_have_one_flag[@]} -ne 0 ]; then
                allflags=("${must_have_one_flag[@]}")
            else
                allflags=("${flags[*]} ${two_word_flags[*]}")
            fi
            COMPREPLY=( $(compgen -W "${allflags[*]}" -- "$cur") )
            if [[ $(type -t compopt) = "builtin" ]]; then
                [[ "${COMPREPLY[0]}" == *= ]] || compopt +o nospace
            fi

            # complete after --flag=abc
            if [[ $cur == *=* ]]; then
                if [[ $(type -t compopt) = "builtin" ]]; then
                    compopt +o nospace
                fi

                local index flag
                flag="${cur%=*}"
                __fuzzit_index_of_word "${flag}" "${flags_with_completion[@]}"
                COMPREPLY=()
                if [[ ${index} -ge 0 ]]; then
                    PREFIX=""
                    cur="${cur#*=}"
                    ${flags_completion[${index}]}
                    if [ -n "${ZSH_VERSION}" ]; then
                        # zsh completion needs --flag= prefix
                        eval "COMPREPLY=( \"\${COMPREPLY[@]/#/${flag}=}\" )"
                    fi
                fi
            fi
            return 0;
            ;;
    esac

    # check if we are handling a flag with special work handling
    local index
    __fuzzit_index_of_word "${prev}" "${flags_with_completion[@]}"
    if [[ ${index} -ge 0 ]]; then
        ${flags_completion[${index}]}
        return
    fi

    # we are parsing a flag and don't have a special handler, no completion
    if [[ ${cur} != "${words[cword]}" ]]; then
        return
    fi

    local completions
    completions=("${commands[@]}")
    if [[ ${#must_have_one_noun[@]} -ne 0 ]]; then
        completions=("${must_have_one_noun[@]}")
    fi
    if [[ ${#must_have_one_flag[@]} -ne 0 ]]; then
        completions+=("${must_have_one_flag[@]}")
    fi
    COMPREPLY=( $(compgen -W "${completions[*]}" -- "$cur") )

    if [[ ${#COMPREPLY[@]} -eq 0 && ${#noun_aliases[@]} -gt 0 && ${#must_have_one_noun[@]} -ne 0 ]]; then
        COMPREPLY=( $(compgen -W "${noun_aliases[*]}" -- "$cur") )
    fi

    if [[ ${#COMPREPLY[@]} -eq 0 ]]; then
		if declare -F __fuzzit_custom_func >/dev/null; then
			# try command name qualified custom func
			__fuzzit_custom_func
		else
			# otherwise fall back to unqualified for compatibility
			declare -F __custom_func >/dev/null && __custom_func
		fi
    fi

    # available in bash-completion >= 2, not always present on macOS
    if declare -F __ltrim_colon_completions >/dev/null; then
        __ltrim_colon_completions "$cur"
    fi

    # If there is only 1 completion and it is a flag with an = it will be completed
    # but we don't want a space after the =
    if [[ "${#COMPREPLY[@]}" -eq "1" ]] && [[ $(type -t compopt) = "builtin" ]] && [[ "${COMPREPLY[0]}" == --*= ]]; then
       compopt -o nospace
    fi
}

# The arguments should be in the form "ext1|ext2|extn"
__fuzzit_handle_filename_extension_flag()
{
    local ext="$1"
    _filedir "@(${ext})"
}

__fuzzit_handle_subdirs_in_dir_flag()
{
    local dir="$1"
    pushd "${dir}" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1
}

__fuzzit_handle_flag()
{
    __fuzzit_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    # if a command required a flag, and we found it, unset must_have_one_flag()
    local flagname=${words[c]}
    local flagvalue
    # if the word contained an =
    if [[ ${words[c]} == *"="* ]]; then
        flagvalue=${flagname#*=} # take in as flagvalue after the =
        flagname=${flagname%=*} # strip everything after the =
        flagname="${flagname}=" # but put the = back
    fi
    __fuzzit_debug "${FUNCNAME[0]}: looking for ${flagname}"
    if __fuzzit_contains_word "${flagname}" "${must_have_one_flag[@]}"; then
        must_have_one_flag=()
    fi

    # if you set a flag which only applies to this command, don't show subcommands
    if __fuzzit_contains_word "${flagname}" "${local_nonpersistent_flags[@]}"; then
      commands=()
    fi

    # keep flag value with flagname as flaghash
    # flaghash variable is an associative array which is only supported in bash > 3.
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        if [ -n "${flagvalue}" ] ; then
            flaghash[${flagname}]=${flagvalue}
        elif [ -n "${words[ $((c+1)) ]}" ] ; then
            flaghash[${flagname}]=${words[ $((c+1)) ]}
        else
            flaghash[${flagname}]="true" # pad "true" for bool flag
        fi
    fi

    # skip the argument to a two word flag
    if [[ ${words[c]} != *"="* ]] && __fuzzit_contains_word "${words[c]}" "${two_word_flags[@]}"; then
			  __fuzzit_debug "${FUNCNAME[0]}: found a flag ${words[c]}, skip the next argument"
        c=$((c+1))
        # if we are looking for a flags value, don't show commands
        if [[ $c -eq $cword ]]; then
            commands=()
        fi
    fi

    c=$((c+1))

}

__fuzzit_handle_noun()
{
    __fuzzit_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    if __fuzzit_contains_word "${words[c]}" "${must_have_one_noun[@]}"; then
        must_have_one_noun=()
    elif __fuzzit_contains_word "${words[c]}" "${noun_aliases[@]}"; then
        must_have_one_noun=()
    fi

    nouns+=("${words[c]}")
    c=$((c+1))
}

__fuzzit_handle_command()
{
    __fuzzit_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    local next_command
    if [[ -n ${last_command} ]]; then
        next_command="_${last_command}_${words[c]//:/__}"
    else
        if [[ $c -eq 0 ]]; then
            next_command="_fuzzit_root_command"
        else
            next_command="_${words[c]//:/__}"
        fi
    fi
    c=$((c+1))
    __fuzzit_debug "${FUNCNAME[0]}: looking for ${next_command}"
    declare -F "$next_command" >/dev/null && $next_command
}

__fuzzit_handle_word()
{
    if [[ $c -ge $cword ]]; then
        __fuzzit_handle_reply
        return
    fi
    __fuzzit_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"
    if [[ "${words[c]}" == -* ]]; then
        __fuzzit_handle_flag
    elif __fuzzit_contains_word "${words[c]}" "${commands[@]}"; then
        __fuzzit_handle_command
    elif [[ $c -eq 0 ]]; then
        __fuzzit_handle_command
    elif __fuzzit_contains_word "${words[c]}" "${command_aliases[@]}"; then
        # aliashash variable is an associative array which is only supported in bash > 3.
        if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
            words[c]=${aliashash[${words[c]}]}
            __fuzzit_handle_command
        else
            __fuzzit_handle_noun
        fi
    else
        __fuzzit_handle_noun
    fi
    __fuzzit_handle_word
}

_fuzzit_auth()
{
    last_command="fuzzit_auth"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_completion()
{
    last_command="fuzzit_completion"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--help")
    flags+=("-h")
    local_nonpersistent_flags+=("--help")
    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_create_job()
{
    last_command="fuzzit_create_job"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--args=")
    two_word_flags+=("--args")
    local_nonpersistent_flags+=("--args=")
    flags+=("--branch=")
    two_word_flags+=("--branch")
    local_nonpersistent_flags+=("--branch=")
    flags+=("--cpus=")
    two_word_flags+=("--cpus")
    local_nonpersistent_flags+=("--cpus=")
    flags+=("--environment=")
    two_word_flags+=("--environment")
    two_word_flags+=("-e")
    local_nonpersistent_flags+=("--environment=")
    flags+=("--revision=")
    two_word_flags+=("--revision")
    local_nonpersistent_flags+=("--revision=")
    flags+=("--type=")
    two_word_flags+=("--type")
    local_nonpersistent_flags+=("--type=")
    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_create_target()
{
    last_command="fuzzit_create_target"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--seed=")
    two_word_flags+=("--seed")
    local_nonpersistent_flags+=("--seed=")
    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_create()
{
    last_command="fuzzit_create"

    command_aliases=()

    commands=()
    commands+=("job")
    commands+=("target")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_download()
{
    last_command="fuzzit_download"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_get()
{
    last_command="fuzzit_get"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_fuzzit_root_command()
{
    last_command="fuzzit"

    command_aliases=()

    commands=()
    commands+=("auth")
    commands+=("completion")
    commands+=("create")
    commands+=("download")
    commands+=("get")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--api-key=")
    two_word_flags+=("--api-key")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

__start_fuzzit()
{
    local cur prev words cword
    declare -A flaghash 2>/dev/null || :
    declare -A aliashash 2>/dev/null || :
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -s || return
    else
        __fuzzit_init_completion -n "=" || return
    fi

    local c=0
    local flags=()
    local two_word_flags=()
    local local_nonpersistent_flags=()
    local flags_with_completion=()
    local flags_completion=()
    local commands=("fuzzit")
    local must_have_one_flag=()
    local must_have_one_noun=()
    local last_command
    local nouns=()

    __fuzzit_handle_word
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_fuzzit fuzzit
else
    complete -o default -o nospace -F __start_fuzzit fuzzit
fi

# ex: ts=4 sw=4 et filetype=sh
