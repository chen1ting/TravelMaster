import {Center, Grid, GridItem, Input, SimpleGrid, Button} from "@chakra-ui/react";
import {useState} from "react";
import {useNavigate} from "react-router-dom";

import ActivityCard from "../../Components/DiscoverComponents/ActivityCard";

//
const ShowActivities = () => {
    let activityDisplayList = []
    for (let i = 1; i < 7; i++) {
        {
            activityDisplayList.push(
                ActivityCard(
                    1,
                    'title 2',
                    4,
                    true,
                    'cat4',
                    'https://bit.ly/2Z4KKcF'
                )
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
    const [searchInput, setSearchInput] = useState('');
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
                            navigate("/createactivity");
                        }}
                        w={200}
                        m={4}
                        bg="teal"
                    >
                        <font size={5} color={"white"}>
                            Create an activity
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
                        onChange={(e) => searchInput(e.target.value)} // Change this
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