import {Center, Container, Group, Text} from "@mantine/core";
import Footer from "@/components/Footer";


const containerStyle = {
    height: "calc(100vh - 80px)"
}

const borderStyle = {
    borderRight: "1px solid rgba(0,0,0,.3)",
    paddingRight: "24px"
}

export default function NotFoundPage() {

    return (
        <>
            <Container fluid className="flex justify-center align-middle" style={containerStyle}>
                <Center className="text-center">
                    <Group>
                        <Text fw={500} fz={24} style={borderStyle}>404</Text>
                        <Text>This page could not be found.</Text>
                    </Group>
                </Center>
            </Container>
            <Footer/>
        </>
    )
}