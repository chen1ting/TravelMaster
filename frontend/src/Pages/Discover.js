import {
    Box, Button,
    Grid,
    GridItem,
    Text
} from "@chakra-ui/react";
import DiscoverActivities from "./DiscoverSubPages/DiscoverActivities";
import {useNavigate} from "react-router-dom";

const Discover = () => {
    const navigate = useNavigate();

    return (
        <Grid
            templateAreas={`"title createNew"
                            "content content"
                            `}
            gridTemplateRows={'10fr 90fr'}
            h='100vh'
            fontWeight='bold'
            bgColor={'blue.50'}
        >
            <GridItem area={"title"}>
                <Box position={'relative'} top={'40%'} left={'17.5%'} transform={'translate(-50%,-50%)'}
                     textAlign={"left"} width='fit-content' mt={'1%'}>
                    <Text fontSize='4xl'>Discover</Text>
                </Box>
            </GridItem>
            <GridItem area={'createNew'}>
                <Box position={'relative'} top={'40%'} left={'80%'} transform={'translate(-50%,-50%)'}
                     textAlign={"left"} width='fit-content' mt={'1%'}>
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
                </Box>
            </GridItem>
            <GridItem area={"content"} bgColor={'white'}>
                <DiscoverActivities></DiscoverActivities>
            </GridItem>
        </Grid>
    );
}

export default Discover;