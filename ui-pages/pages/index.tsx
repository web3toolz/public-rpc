import * as React from "react"
import {useState} from "react"
import {useDebouncedValue} from "@mantine/hooks";
import {Center, Container, Title} from "@mantine/core";

import RpcCardsGrid from "../components/RpcCardGrid/RpcCardsGrid";
import SearchBar from "../components/SearchBar";
import {Rpc} from "@/models/rpc"
import {useFilterRpcData} from "@/hooks/filterRpcData";
import {fetchRpcData} from "@/api/fetchRpcData";

const pageStyles = {
    backgroundColor: "rgb(245 245 245/var(--tw-bg-opacity))",
}

interface HomeProps {
    data: Rpc[],
    error: any,
}

export async function getStaticProps(): Promise<{ props: HomeProps }> {
    let data: Rpc[] = [];
    let error: Error | null = null;

    try {
        data = await fetchRpcData();
    } catch (e: any) {
        error = e;
    }

    return {
        props: {
            data: data,
            error: error
        },
    }
}

export default function Home({data, error}: HomeProps) {
    const [query, setQuery] = useState<string>("");
    const [queryDebounced] = useDebouncedValue(query, 200);

    const filteredData: Rpc[] = useFilterRpcData(data, queryDebounced);
    const titleText: string = "Find free RPC endpoint for any blockchain";

    return (
        <Container fluid h={100} style={pageStyles}>
            <Center className="my-10 text-center">
                <Title order={1}>{titleText}</Title>
            </Center>
            <SearchBar query={query} setQuery={setQuery}/>
            <RpcCardsGrid rpcData={filteredData}/>
        </Container>
    )
}
