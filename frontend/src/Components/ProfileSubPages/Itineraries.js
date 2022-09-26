import React, {useState} from "react";
import "react-datepicker/dist/react-datepicker.css";
import "./customDatePickerWidth.css"
import {Box, Button, Center, Grid, GridItem} from "@chakra-ui/react";
import Calendar from "react-calendar";
import 'react-calendar/dist/Calendar.css';


const Itineraries = () => {
    const mark = ['22-09-2022', '03-12-2022', '05-12-2022'];
    const [value, onChange] = useState(new Date());

    return (
        <Box>
            <Box mb={3}>
                <Center>
                    <Calendar
                        onChange={onChange}
                        value={value}
                        minDate={new Date(2010, 1, 1)}
                        showDoubleView={true}
                        minDetail={'year'}
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