import * as React from "react";
import {Button, Card, ColorSwatch, CopyButton, Group, Text, TextInput, Tooltip} from '@mantine/core';
import {Rpc} from "@/models/rpc"
import {capitalize, extractHostname} from "@/utils";

interface RpcCardProps {
    rpc: Rpc;
}

const statusColorMap: StringMap = {
    active: "#7CFC00",
    inactive: "red",
}

const statusLabelMap: StringMap = {
    active: "Active",
    inactive: "Inactive",
}


function RpcCard({rpc}: RpcCardProps): React.ReactElement {
    const rpcUrl = rpc.http || rpc.ws;
    const title = extractHostname(rpcUrl);
    const chain = capitalize(rpc.chain);
    const network = capitalize(rpc.network);
    const status = rpc.status;

    const statusColor: string | undefined = statusColorMap[status];
    const statusLabel: string | undefined = statusLabelMap[status];


    return (
        <Card shadow="sm" padding="md" radius="md" withBorder>
            <Group justify="space-between">
                <Text fw={600}>{title}</Text>
                {statusColor &&
                    <Tooltip label={statusLabel}>
                        <ColorSwatch color={statusColor} size={18}/>
                    </Tooltip>
                }
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