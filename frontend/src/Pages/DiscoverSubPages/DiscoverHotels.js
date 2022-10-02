import {
    Box,
    Button,
    Flex,
    Grid,
    GridItem,
    Slider,
    SliderFilledTrack,
    SliderThumb,
    SliderTrack,
    Text
} from "@chakra-ui/react";
import React, {useState} from "react";
import {
    NumberInput,
    NumberInputField,
    NumberInputStepper,
    NumberIncrementStepper,
    NumberDecrementStepper,
} from '@chakra-ui/react'
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./DatePickerWidth.css"

const DiscoverHotels = () => {
    const [dateRange, setDateRange] = useState([null, null]);
    const [checkIn, checkOut] = dateRange;
    const [value, setValue] = React.useState(2)
    const handleChange = (value) => setValue(value)

    function onSubmit(e) {
        e.preventDefault();
    }

    return (
        <Grid
            templateAreas={`"inputDates inputTravellers"
                            "results results"
                            "loadMore loadMore"
                            `}
            gridTemplateRows={'1fr 8fr 1fr'}
            gridTemplateColumns={'1fr 1fr'}
            h='100vh'
        >
            <GridItem area={'inputDates'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"} pl={'35%'} pr={'15%'}>
                    <Text fontSize={'xl'} mb={2}>Select Check-in & Check-out Dates</Text>
                    <Box className={"customDatePickerWidth"} padding={'1px'} bgColor={"black"}>
                        <DatePicker
                            selectsRange={true}
                            startDate={checkIn}
                            endDate={checkOut}
                            onChange={(update) => {
                                setDateRange(update);
                            }}
                            monthsShown={2}
                            isClearable={true}
                            showMonthDropdown
                            showYearDropdown
                            dropdownMode="select"
                            dateFormat="dd MMM yyyy"
                        />
                    </Box>
                </Box>
            </GridItem>
            <GridItem area={'inputTravellers'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"} pl={'20%'} pr={'20%'}>
                    <Text fontSize={'xl'} mb={2}>Number of Travellers</Text>
                    <Flex>
                        <NumberInput defaultValue={2} min={1} max={14} maxW='100px' mr='2rem' value={value}
                                     onChange={handleChange}>
                            <NumberInputField/>
                            <NumberInputStepper>
                                <NumberIncrementStepper/>
                                <NumberDecrementStepper/>
                            </NumberInputStepper>
                        </NumberInput>
                        <Slider
                            flex='1'
                            focusThumbOnChange={false}
                            value={value}
                            onChange={handleChange}
                            defaultValue={2}
                            min={1}
                            max={14}
                        >
                            <SliderTrack>
                                <SliderFilledTrack/>
                            </SliderTrack>
                            <SliderThumb fontSize='md' boxSize='32px' children={value}/>
                        </Slider>
                    </Flex>
                </Box>
            </GridItem>
            <GridItem area={"results"} pt={2}>
                test
            </GridItem>
            <GridItem area={'loadMore'} bgColor={'white'}>
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

export default DiscoverHotels;