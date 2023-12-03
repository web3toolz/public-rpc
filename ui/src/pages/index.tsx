import * as React from "react"
import {useState} from "react"
import type {PageProps} from "gatsby"
import {useDebouncedValue} from "@mantine/hooks";
import {Title, Center, Container} from "@mantine/core";

import RpcCardsGrid from "../components/RpcCardGrid/RpcCardsGrid";
import SearchBar from "../components/SearchBar";
import Navbar from "../components/Navbar/index";
import Footer from "../components/Footer/index";
import {Rpc} from "../models/rpc"

import {useFetchRpcData} from "../hooks/fetchRpc";
import {useFilterRpcData} from "../hooks/filterRpc";

const pageStyles = {
    backgroundColor: "rgb(245 245 245/var(--tw-bg-opacity))",
}

const IndexPage: React.FC<PageProps> = () => {
    const [query, setQuery] = useState<string>("");
    const [queryDebounced] = useDebouncedValue(query, 200);
    const {data, loading, error,} = useFetchRpcData();
    const filteredData: Rpc[] = useFilterRpcData(data, queryDebounced);
    const titleText: string = "Find free RPC endpoint for any blockchain";


    return (
        <Container fluid h={100} style={pageStyles}>
            {/*<Navbar/>*/}
            <Center className="my-10 text-center">
                <Title order={1}>{titleText}</Title>
            </Center>
            <SearchBar query={query} setQuery={setQuery}/>
            <RpcCardsGrid rpcData={filteredData}/>
            {/*<Footer/>*/}
        </Container>
    )
}

export default IndexPage
