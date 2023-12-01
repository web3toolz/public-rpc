import * as React from "react";
import RpcCard from "./RpcCard";
import {Grid} from '@mantine/core';
import {Rpc} from "../hooks/fetchRpc";


interface RpcCardGridProps {
    rpcData: Rpc[];
}

function RpcCardGrid({rpcData}: RpcCardGridProps): React.ReactElement {
    const cards: React.ReactElement[] = rpcData && rpcData.map((rpc: Rpc) => {
        return <Grid.Col span={{base: 12, md: 4, lg: 3}}><RpcCard rpc={rpc}/></Grid.Col>
    })

    return (
        <Grid className="px-10">
            { cards.length > 0 ? cards : <p>No RPCs found</p>}
        </Grid>
    )
}

export default RpcCardGrid;