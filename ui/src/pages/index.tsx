import * as React from "react"
import {useState} from "react"
import type {HeadFC, PageProps} from "gatsby"
import {useDebouncedValue} from "@mantine/hooks";

import RpcCardsGrid from "../components/RpcCardsGrid";
import SearchBar from "../components/SearchBar";
import Navbar from "../components/Navbar/index";
import Footer from "../components/Footer/index";
import {Rpc} from "../models/rpc"

import {useFetchRpcData} from "../hooks/fetchRpc";
import {useFilterRpcData} from "../hooks/filterRpc";

const pageStyles = {
    backgroundColor: "rgb(245 245 245/var(--tw-bg-opacity))",
    fontFamily: "-apple-system, Roboto, sans-serif, serif",
}

const IndexPage: React.FC<PageProps> = () => {
    const [query, setQuery] = useState<string>("");
    const [queryDebounced] = useDebouncedValue(query, 200);
    const {data, loading, error,} = useFetchRpcData();
    const filteredData: Rpc[] = useFilterRpcData(data, queryDebounced);

    return (
        <main style={pageStyles}>
            <Navbar/>
            <SearchBar query={query} setQuery={setQuery}/>
            <RpcCardsGrid rpcData={filteredData}/>
            <Footer/>
        </main>
    )
}

export default IndexPage

export const Head: HeadFC = () => <title>Public RPC | Web3toolz</title>
