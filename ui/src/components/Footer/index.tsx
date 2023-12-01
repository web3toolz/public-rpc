import * as React from "react";
import {PageProps} from "gatsby";
import {Anchor, Container, Group} from '@mantine/core';
import * as classes from "./footer.module.css";

const links = [
    {link: '#', label: 'Contact'},
    {link: '#', label: 'Privacy'},
    {link: '#', label: 'Blog'},
    {link: '#', label: 'Careers'},
];

const Footer: React.FC<PageProps> = () => {
    const items = links.map((link) => (
        <Anchor<'a'>
            c="dimmed"
            key={link.label}
            href={link.link}
            onClick={(event) => event.preventDefault()}
            size="sm"
        >
            {link.label}
        </Anchor>
    ));

    return (
        <div className={classes.footer}>
            <Container className={classes.inner}>
                <Group className={classes.links}>{items}</Group>
            </Container>
        </div>
    )
}

export default Footer;