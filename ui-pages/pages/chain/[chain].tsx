import {Center, Container, Title} from "@mantine/core";
import Footer from "@/components/Footer";
import Head from "next/head";
import * as React from "react";
import {capitalize} from "@/utils";
import {Rpc} from "@/models/rpc";
import {fetchRpcData} from "@/api/fetchRpcData";
import {useFetchRpcData} from "@/hooks/fetchRpcData";
import RpcCardsGrid from "@/components/RpcCardGrid/RpcCardsGrid";

interface ChainPageProps {
    chain?: string,
    data: Rpc[],
    error: any,
}

function getMetadata(chain: string): React.ReactElement {
    chain = capitalize(chain);
    const title: string = `Free RPC endpoint for ${chain} | Web3toolz.com`;
    const description: string = `Free RPC endpoint for ${chain} chain. Find free RPC endpoint for ${chain} chain`;
    const keywords: string = `${chain}, RPC, EVM, Blockchain, Gas Price, Web3, Ethereum, BNB, Polygon, Optimism, Avalanche`;

    return (
        <Head>
            <title>{title}</title>
            <meta name="description" content={description}/>
            <meta name="keywords" content={keywords}/>
        </Head>
    )
}

export async function getStaticPaths() {
    let data: Rpc[] = [];
    let error: Error | null = null;

    try {
        data = await fetchRpcData({});
    } catch (e: any) {
        error = e;
    }

    if (error) {
        return {
            paths: [],
            fallback: false
        }
    }

    const paths = data.map(rpc => {
        return {params: {chain: rpc.chain}}
    })
    return {paths, fallback: false}
}

export async function getStaticProps({params}: any): Promise<{ props?: ChainPageProps, notFound?: boolean }> {
    let data: Rpc[] = [];
    let error: Error | null = null;
    const chain = params?.chain;

    if (!chain) {
        return {
            notFound: true,
        }
    }

    try {
        data = await fetchRpcData({chain});
    } catch (e: any) {
        error = e;
    }

    if (data.length === 0) {
        return {
            notFound: true,
        }
    }

    return {
        props: {
            data: data,
            error: error,
            chain: chain,
        },
    }
}

export default function ChainPage({data, error, chain}: ChainPageProps) {
    const {rpcData} = useFetchRpcData({initialState: data, chain});

    const titleText = `Free RPC endpoint for ${capitalize(chain || "")} chain`;

    return (
        <>
            {getMetadata(chain || "")}
            <Container fluid h={100}>
                <Center className="my-10 text-center">
                    <Title order={1}>{titleText}</Title>
                </Center>
                <RpcCardsGrid rpcData={rpcData}/>
                <Footer/>
            </Container>
        </>
    )
}