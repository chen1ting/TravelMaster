import {Box, Button, Grid, GridItem, Input, SimpleGrid, Text} from "@chakra-ui/react";
import ActivityCard from "../../Components/ActivityCard";
import {useState} from "react";
import { useNavigate } from "react-router-dom";

const DiscoverActivities = () => {
    const [searchInput] = useState('');
    const navigate = useNavigate();
    function onSubmit(e) {
        e.preventDefault();
    }

    return (
        <Grid
            templateAreas={`"search"
                            "results"
                            "load_more"
                            `}
            gridTemplateRows={'10fr 80fr 10fr'}
            h='100vh'
        >
            <GridItem area={'search'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    <Button
                        as="a"
                        onClick={() => {
                            navigate("/createevent");
                        }}
                        w={200}
                        m={4}
                        bg="teal"
                    >
                        <font size={5} color={"white"}>
                            Create an event
                        </font>
                    </Button>
                    <Input
                        m={4}
                        w={'50%'}
                        size='lg'
                        bgColor={'whitesmoke'}
                        borderColor={'blackAlpha.700'}
                        variant='outline'
                        focusBorderColor='lime'
                        placeholder='Search for events and activities!'
                        onChange={(e) => searchInput(e.target.value)}
                    />
                </Box>
            </GridItem>
            <GridItem area={"results"} bgColor={'white'}>
                <SimpleGrid minChildWidth='300px' spacing='30px' mt={'1%'} ml={'5%'} mr={'5%'}>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                    <ActivityCard></ActivityCard>
                </SimpleGrid>
            </GridItem>
            <GridItem area={'load_more'} bgColor={'white'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"} mb={'5%'} height={"fit-content"}>
                    <Button colorScheme='teal' variant='solid'>
                        View More Results
                    </Button>
                    {/*On click loading more results*/}
                    <Button
                        isLoading
                        loadingText='Loading...'
                        colorScheme='teal'
                        variant='outline'
                    >
                        {/*All results loaded*/}
                        View More Results
                    </Button>
                    <Text fontSize='4xl'>There are no more results to load</Text>
                </Box>
            </GridItem>
        </Grid>
    )
}

export default DiscoverActivities;