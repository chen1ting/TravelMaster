import {Center, Grid, GridItem, Input, SimpleGrid, Button} from "@chakra-ui/react";
import {useState} from "react";
import { useNavigate } from "react-router-dom";

import ActivityCard from "../../Components/DiscoverComponents/ActivityCard";

const ShowActivities = () => {
    let activityDisplayList = []
    for (let i = 0; i < 2; i++) {
        {
            activityDisplayList.push(
                ActivityCard(0,'title 1', 3, false, 'cat1', 'https://bit.ly/2Z4KKcF')
            )
        }
    }
    for (let i = 5; i < 7; i++) {
        {
            activityDisplayList.push(
                ActivityCard(1,'title 2', 4, true, 'cat4', 'https://bit.ly/2Z4KKcF')
            )
        }
    }
    return (
        <SimpleGrid minChildWidth='300px' spacing='30px' mt={'1%'} ml={'5%'} mr={'5%'}>
            {activityDisplayList}
        </SimpleGrid>
    )
}


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
                            `}
            gridTemplateRows={'10fr 90fr'}
            h='80vh'
        >
            <GridItem area={'search'}>


                <Center>
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
                        placeholder='Search for activities!'
                        onChange={(e) => searchInput(e.target.value)}
                    />
                </Center>
            </GridItem>
            <GridItem area={"results"} bgColor={'white'}>
                {ShowActivities()}
            </GridItem>
        </Grid>
    )
}

export default DiscoverActivities;