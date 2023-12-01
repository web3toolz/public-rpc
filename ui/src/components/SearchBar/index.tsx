import * as React from 'react';
import {useState} from 'react';
import {TextInput} from '@mantine/core';
import {PageProps} from "gatsby";
import * as classes from './searchbar.module.css';


const SearchBar: React.FC<PageProps> = () => {
    const [focused, setFocused] = useState(false);
    const [value, setValue] = useState('');
    const floating = value.trim().length !== 0 || focused || undefined;

    return (
        <div className="flex justify-center p-10 my-14">
            <TextInput
                label="Search for RPCs"
                placeholder="Search for RPCs"
                className="w-screen lg:w-1/3"
                classNames={classes}
                value={value}
                onChange={(event) => setValue(event.currentTarget.value)}
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