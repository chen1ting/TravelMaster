import {
    Box,
    Grid,
    GridItem,
    Text
} from "@chakra-ui/react";
import DiscoverActivities from "./DiscoverSubPages/DiscoverActivities";

const Discover = () => {
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
            <GridItem area={"title"}>
                <Box position={'relative'} top={'23%'} left={'7.5%'} transform={'translate(-50%,-50%)'}
                     textAlign={"left"} width='fit-content' mt={'1%'}>
                    <Text fontSize='4xl'>Discover</Text>
                </Box>
            </GridItem>
            <GridItem area={"content"} bgColor={'white'}>
                <DiscoverActivities></DiscoverActivities>
            </GridItem>
        </Grid>
    );
}

export default Discover;