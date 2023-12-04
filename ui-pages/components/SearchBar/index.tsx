import * as React from 'react';
import {useState} from 'react';
import {TextInput} from '@mantine/core';
import "./searchbar.module.css"

interface SearchBarProps {
    query: string;
    setQuery: (query: string) => void;
}

function SearchBar({query, setQuery}: SearchBarProps): React.ReactElement {
    const [focused, setFocused] = useState(false);
    const floating = (query || "").trim().length !== 0 || focused || undefined;

    return (
        <div className="flex justify-center mt-10 mb-20 px-5">
            <TextInput
                size="md"
                radius="lg"
                placeholder="Search for RPCs"
                className="root w-screen lg:w-1/3"
                value={query}
                onChange={(event) => setQuery(event.currentTarget.value)}
                onFocus={() => setFocused(true)}
                onBlur={() => setFocused(false)}
                mt="md"
                autoComplete="nope"
                data-floating={floating}
                labelProps={{'data-floating': floating}}
            />
        </div>

    )
}

export default SearchBar