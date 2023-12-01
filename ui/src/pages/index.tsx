import * as React from "react"
import type { HeadFC, PageProps } from "gatsby"

import RpcCardsGrid from "../components/RpcCardsGrid";
import SearchBar from "../components/SearchBar";
import Navbar from "../components/Navbar/index";
import Footer from "../components/Footer/index";

import {useFetchRpcData, Rpc} from "../hooks/fetchRpc";

const pageStyles = {
  backgroundColor: "rgb(245 245 245/var(--tw-bg-opacity))",
  fontFamily: "-apple-system, Roboto, sans-serif, serif",
}

const IndexPage: React.FC<PageProps> = () => {
    const {data, loading, error, } = useFetchRpcData();

  return (
    <main style={pageStyles}>
        <Navbar/>
        <SearchBar />
        <RpcCardsGrid rpcData={data} />
        <Footer />
    </main>
  )
}

export default IndexPage

export const Head: HeadFC = () => <title>Public RPC | Web3toolz</title>
