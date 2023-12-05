import React from 'react';
import {ActionIcon, Container, Group, rem, Text, Anchor} from '@mantine/core';
import {IconBrandTwitter} from '@tabler/icons-react';
import classes from './footer.module.css'


const twitterUrl: string = "https://twitter.com/web3toolz";

function Footer(): React.ReactElement {


    return (
        <div className={classes.footer}>
            <Container className={classes.inner}>
                <Text className={classes.text}>© 2023 web3toolz.com</Text>
                <Group className={classes.links} gap="xs" justify="flex-end" wrap="nowrap">
                    <Anchor href={twitterUrl}>
                        <ActionIcon size="lg" color="black" variant="subtle" >
                            <IconBrandTwitter style={{width: rem(18), height: rem(18)}} stroke={1.5}/>
                        </ActionIcon>
                    </Anchor>
                </Group>
            </Container>
        </div>
    );
}

export default Footer;