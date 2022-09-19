import {
    Box, Button,
    Grid,
    GridItem,
    Input,
    SimpleGrid,
    Text
} from "@chakra-ui/react";
import {useState} from "react";
import ActivityCard from '../Components/ActivityCard'

const Discover = () => {
    const [searchInput] = useState('');

    function onSubmit(e) {
        e.preventDefault();
    }

    return (
        <Grid
            templateAreas={`"title"
                            "search"
                            "results"
                            "load_more"
                            `}
            gridTemplateRows={'10fr 10fr 80fr'}
            h='100vh'
            gap='2'
            fontWeight='bold'
            bgColor={'blue.50'}
        >
            <GridItem area={"title"}>
                <Box position={'relative'} top={'50%'} left={'7.5%'} transform={'translate(-50%,-50%)'}
                     textAlign={"left"} width='fit-content' mt={'1%'}>
                    <Text fontSize='4xl'>Discover</Text>
                </Box>
            </GridItem>
            <GridItem area={'search'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
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
    );
}

export default Discover;