import {SimpleGrid} from "@chakra-ui/react";
import itineraryCard from "./ItineraryCard";

const DisplayItineraries = () => {
    let itineraryDisplayList = []
    for (let i = 0; i < 3; i++) {
        {
            itineraryDisplayList.push(
                itineraryCard(
                    1,
                    'Itinerary 2',
                    (new Date()).toLocaleDateString(),
                    (new Date()).toLocaleDateString(),
                )
            )
        }
    }
    return (
        <SimpleGrid minChildWidth='300px' spacing='30px' mt={'1%'} ml={'5%'} mr={'5%'}>
            {itineraryDisplayList}
        </SimpleGrid>
    )
}

export default DisplayItineraries;