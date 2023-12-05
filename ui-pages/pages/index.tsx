import * as React from "react"
import {useState} from "react"
import {useDebouncedValue} from "@mantine/hooks";
import {Center, Container, Title} from "@mantine/core";

import RpcCardsGrid from "@/components/RpcCardGrid/RpcCardsGrid";
import SearchBar from "@/components/SearchBar";
import Footer from "@/components/Footer";
import {Rpc} from "@/models/rpc"
import {useFilterRpcData} from "@/hooks/filterRpcData";
import {useFetchRpcData} from "@/hooks/fetchRpcData";
import {fetchRpcData} from "@/api/fetchRpcData";

const titleText: string = "Find free RPC endpoint for any EVM and non-EVM chain";

interface HomeProps {
    data: Rpc[],
    error: any,
}

export async function getStaticProps(): Promise<{ props: HomeProps }> {
    let data: Rpc[] = [];
    let error: Error | null = null;

    try {
        data = await fetchRpcData({});
    } catch (e: any) {
        error = e;
    }

    return {
        props: {
            data: data.slice(0, 10),
            error: error
        },
    }
}

export default function Home({data, error}: HomeProps) {
    const {rpcData} = useFetchRpcData({initialState: data});
    const [query, setQuery] = useState<string>("");
    const [queryDebounced] = useDebouncedValue(query, 200);

    const filteredData: Rpc[] = useFilterRpcData(rpcData, queryDebounced);


    return (
        <Container fluid h={100}>
            <Center className="my-10 text-center">
                <Title order={1}>{titleText}</Title>
            </Center>
            <SearchBar query={query} setQuery={setQuery}/>
            <RpcCardsGrid rpcData={filteredData}/>
            <Footer/>
        </Container>
    )
}
