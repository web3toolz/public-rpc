import * as React from "react";
import { useState, useEffect } from 'react';
import RpcCard from "./RpcCard";
import {Group, Pagination, SimpleGrid} from '@mantine/core';
import {Rpc} from "../models/rpc";
import {serialize} from "v8";


interface RpcCardGridProps {
    rpcData: Rpc[];
}

const gridParams = {base: 1, sm: 1, md: 2, lg: 4};
const itemsPerPage: number = 20;

function RpcCardGrid({rpcData}: RpcCardGridProps): React.ReactElement {
    const [activePage, setPage] = useState<number>(1);
    const totalPages: number = Math.ceil(rpcData.length / itemsPerPage);
    const sliceStart: number = (activePage - 1) * itemsPerPage;
    const sliceEnd: number = activePage * itemsPerPage;

    const cards: React.ReactElement[] = rpcData && rpcData.map((rpc: Rpc) => {
        return <RpcCard rpc={rpc}/>
    }).slice(sliceStart, sliceEnd);

    useEffect(() => {
        setPage(1);
    }, [rpcData])

    return (
        <>
            <SimpleGrid className="px-10 mb-20" cols={gridParams} spacing="lg">
                {cards.length > 0 ? cards : <p>No RPCs found</p>}
            </SimpleGrid>
            <Group justify="center">
                <Pagination size="lg" total={totalPages} value={activePage} onChange={setPage}/>
            </Group>
        </>
    )
}

export default RpcCardGrid;