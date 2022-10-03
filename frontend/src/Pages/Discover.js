import {
    Button, Center, Flex,
    Grid,
    GridItem, Spacer,
} from "@chakra-ui/react";
import DiscoverActivities from "./DiscoverSubPages/DiscoverActivities";
import {useNavigate} from "react-router-dom";

const Discover = () => {
    const navigate = useNavigate();

    return (
        <Grid
            templateAreas={`"title"
                            "content"
                            `}
            gridTemplateRows={'10fr 90fr'}
            h='100vh'
            fontWeight='bold'
            bgColor={'blue.50'}
        >
            <GridItem area={"title"} mt={3} ml={'4%'} mr={'4%'}>
                <Flex>
                    <Center fontSize='4xl'>Discover</Center>
                    <Spacer />
                    <Button
                        as="a"
                        onClick={() => {
                            navigate("/createevent");
                        }}
                        w={250}
                        m={4}
                        bg="teal"
                    >
                        <font size={5} color={"white"}>
                            Create New Activity
                        </font>
                    </Button>
                </Flex>
            </GridItem>

            <GridItem area={"content"} bgColor={'white'}>
                <DiscoverActivities></DiscoverActivities>
            </GridItem>
        </Grid>
    );
}

export default Discover;