import React, { useState } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./customDatePickerWidth.css"
import {Box, Grid, GridItem, Text} from "@chakra-ui/react";
// import 'react-datepicker/dist/react-datepicker-cssmodules.css';

const DatesSelector = () => {
    const [dateRange, setDateRange] = useState([null, null]);
    const [startDate, endDate] = dateRange;
    return (
        <DatePicker
            selectsRange={true}
            startDate={startDate}
            endDate={endDate}
            onChange={(update) => {
                setDateRange(update);
            }}
            monthsShown={3}
            isClearable={true}
            showMonthDropdown
            showYearDropdown
            dropdownMode="select"
            dateFormat="dd MMM yyyy"
        />
    );
};


const Itineraries = () => {
    return (
        <Grid
            templateAreas={`"datepicker"
                            "content"`}
            gridTemplateRows={'15% 85%'}
            h={'74.5vh'}
        >
            <GridItem area={'datepicker'}>
                <Text>Select Date</Text>
                <Box className={"customDatePickerWidth"} padding={'1px'} bgColor={"black"}>
                    <DatesSelector></DatesSelector>
                </Box>
            </GridItem>
            <GridItem area={'content'} bgColor={'blue.50'}>
                test
            </GridItem>
        </Grid>
    )
}

export default Itineraries;