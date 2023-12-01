import * as React from "react";
import {Button, Card, CopyButton, TextInput, Group, Text, Image} from '@mantine/core';
import {Rpc} from "../models/rpc"
import {capitalize, extractHostname} from "../utils";


interface RpcCardProps {
    rpc: Rpc;
}

function RpcCard({rpc}: RpcCardProps): React.ReactElement {
    const rpcUrl = rpc.http || rpc.ws;
    const title = extractHostname(rpcUrl);
    const chain = capitalize(rpc.chain);
    const network = capitalize(rpc.network);


    return (
        <Card shadow="sm" padding="md" radius="md" withBorder>
            <Group justify="space-between">
                <Text fw={500}>{title}</Text>
                <Text>{rpc.status}</Text>
            </Group>
            <Group className="mt-4">
                <Text>{chain} {network}</Text>
            </Group>
            <TextInput className="mt-4" value={rpcUrl} variant="filled"/>
            <CopyButton value={rpcUrl}>
                {({copied, copy}) => (
                    <Button className="mt-2" color={copied ? 'teal' : 'blue'} onClick={copy}>
                        {copied ? 'Copied' : 'Copy URL'}
                    </Button>
                )}
            </CopyButton>
        </Card>
    )
}

export default RpcCard;