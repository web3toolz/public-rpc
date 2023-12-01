import * as React from 'react';
import {useState} from 'react';
import {TextInput} from '@mantine/core';
import * as classes from './searchbar.module.css';

interface SearchBarProps {
    query: string;
    setQuery: (query: string) => void;
}

function SearchBar({query, setQuery}: SearchBarProps): React.ReactElement {
    const [focused, setFocused] = useState(false);
    const floating = (query || "").trim().length !== 0 || focused || undefined;

    return (
        <div className="flex justify-center p-10 my-14">
            <TextInput
                label="Search for RPCs"
                placeholder="Search for RPCs"
                className="w-screen lg:w-1/3"
                classNames={classes}
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