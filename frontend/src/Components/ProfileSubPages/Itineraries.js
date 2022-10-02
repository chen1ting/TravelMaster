// import "react-datepicker/dist/react-datepicker.css";
// import "./customDatePickerWidth.css"
import {Button, Center, Flex, Grid, GridItem, Spacer} from "@chakra-ui/react";
// import Calendar from "react-calendar";
// import 'react-calendar/dist/Calendar.css';
import {useNavigate} from "react-router-dom";
import DisplayItineraries from "../ItinerariesComponents/DisplayItineraries";


const Itineraries = () => {
    const navigate = useNavigate();
    // const [value, onChange] = useState(new Date());

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
                    <Center fontSize='4xl'>Itineraries</Center>
                    <Spacer />
                    <Button
                        as="a"
                        onClick={() => {
                            navigate("/createitinerary");
                        }}
                        w={250}
                        m={4}
                        bg="teal"
                    >
                        <font size={5} color={"white"}>
                            Create New Itinerary
                        </font>
                    </Button>
                </Flex>
            </GridItem>

            <GridItem area={"content"} bgColor={'white'}>
                <DisplayItineraries></DisplayItineraries>
            </GridItem>
        </Grid>


        // <Box>
        //     <Box mt={3} mb={3}>
        //         <Center>
        //             <Calendar
        //                 onChange={onChange}
        //                 value={value}
        //                 minDate={new Date(2010, 1, 1)}
        //                 showDoubleView={true}
        //                 minDetail={'year'}
        //                 showNeighboringMonth={false}
        //             />
        //         </Center>
        //     </Box>
        //     <Box bgColor={'blue.50'}>
        //         <Button>
        //             Create New Itinerary
        //         </Button>
        //     </Box>
        // </Box>
    );
}

export default Itineraries;