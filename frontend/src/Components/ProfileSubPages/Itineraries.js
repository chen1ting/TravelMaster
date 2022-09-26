import React, {useState} from "react";
import "react-datepicker/dist/react-datepicker.css";
import "./customDatePickerWidth.css"
import {Box, Button, Center} from "@chakra-ui/react";
import Calendar from "react-calendar";
import 'react-calendar/dist/Calendar.css';


const Itineraries = () => {
    const [value, onChange] = useState(new Date());

    return (
        <Box>
            <Box mt={3} mb={3}>
                <Center>
                    <Calendar
                        onChange={onChange}
                        value={value}
                        minDate={new Date(2010, 1, 1)}
                        showDoubleView={true}
                        minDetail={'year'}
                        showNeighboringMonth={false}
                    />
                </Center>
            </Box>
            <Box bgColor={'blue.50'}>
                <Button>
                    Create New Itinerary
                </Button>
            </Box>
        </Box>
    )
}

export default Itineraries;