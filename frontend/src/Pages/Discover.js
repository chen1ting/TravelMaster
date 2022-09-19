import {
    Box,
    Grid,
    GridItem,
    Input,
    InputGroup,
    InputLeftElement,
    InputRightElement,
    SimpleGrid,
    Text
} from "@chakra-ui/react";
import {useState} from "react";
import ActivityCard from '../Components/ActivityCard'
import {PhoneIcon, Search2Icon} from "@chakra-ui/icons";

const Discover = () => {
    const [searchInput] = useState('');

    function onSubmit(e) {
        e.preventDefault();
        //signIn({ username, password });
    }

    return (
        <Grid
            templateAreas={`"row_1"
                            "row_2"
                            "row_3"
                            `}
            gridTemplateRows={'10fr 10fr 80fr'}
            // gridTemplateColumns={'3fr 2fr'}
            h='100vh'
            gap='2'
            // color='blackAlpha.700'
            fontWeight='bold'
        >
            <GridItem area={"row_1"}>
                <Box position={'relative'} top={'50%'} left={'7.5%'} transform={'translate(-50%,-50%)'}
                     textAlign={"left"} width='fit-content'>
                    <Text fontSize='4xl'>Discover</Text>
                </Box>
            </GridItem>
            <GridItem area={'row_2'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    {/*<InputGroup>*/}
                    {/*    <InputLeftElement*/}
                    {/*        children={<Search2Icon color='black'/>}*/}
                    {/*    />*/}
                    {/*    <Input*/}
                    {/*        m={4}*/}
                    {/*        w={'50%'}*/}
                    {/*        bgColor={'whitesmoke'}*/}
                    {/*        variant='outline'*/}
                    {/*        placeholder='Search for events and activities!'*/}
                    {/*        onChange={(e) => searchInput(e.target.value)}*/}
                    {/*    />*/}
                    {/*</InputGroup>*/}
                    <Input
                        m={4}
                        w={'50%'}
                        bgColor={'whitesmoke'}
                        variant='outline'
                        placeholder='Search for events and activities!'
                        onChange={(e) => searchInput(e.target.value)}
                    />
                </Box>
            </GridItem>
            <GridItem area={"row_3"} bgColor={'blue.50'}>
                <SimpleGrid minChildWidth='300px' spacing='30px' ml={'5%'} mr={'5%'}>
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
        </Grid>
    );
}

export default Discover;