import {
    Box,
    Button,
    Flex, FormControl, FormLabel,
    Grid,
    GridItem, Input, Select, SimpleGrid,
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
import FlightCard from "../../Components/DiscoverComponents/FlightCard";

const DiscoverFlights = () => {
    const [location, setLocation] = useState(['', ''])
    const [dateRange, setDateRange] = useState([null, null]);
    const [departure, arrival] = dateRange;
    const [value, setValue] = React.useState(2)
    const handleChange = (value) => setValue(value)

    function onSubmit(e) {
        e.preventDefault();
    }

    return (
        <Grid
            templateAreas={`"inputLocation inputDates inputTravellers ticketTier"
                            "results results results results"
                            "loadMore loadMore loadMore loadMore"
                            `}
            gridTemplateRows={'1fr 6fr 3fr'}
            gridTemplateColumns={'28fr 28fr 28fr 16fr'}
            h='100vh'
        >
            <GridItem area={'inputLocation'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    <Input
                        m={4}
                        w={'80%'}
                        size='lg'
                        bgColor={'white'}
                        borderColor={'blackAlpha.700'}
                        variant='outline'
                        focusBorderColor='lime'
                        placeholder='Leaving from'
                        onChange={(e) => setLocation(e.target.value)}
                    />
                </Box>
            </GridItem>
            <GridItem area={'inputDates'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"} pl={'5%'}>
                    <Text fontSize={'xl'} mb={'2'}>Select Departure and Return Dates</Text>
                    <Box className={"customDatePickerWidth"} padding={'1px'} bgColor={"black"}>
                        <DatePicker
                            selectsRange={true}
                            startDate={departure}
                            endDate={arrival}
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
                     textAlign={"center"} pl={'18%'} pr={'18%'}>
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
            <GridItem area={'ticketTier'}>
                <FormControl isRequired pr={'10%'}>
                    <FormLabel textAlign={'center'} fontSize={'xl'} fontWeight={'bold'}>Ticket</FormLabel>
                    <Select placeholder='Select Ticket Class'>
                        <option value='Economy'>Economy</option>
                        <option value='Premium Economy'>Premium Economy</option>
                        <option value='Business CLass'>Business Class</option>
                        <option value='First CLass'>First Class</option>
                    </Select>
                </FormControl>
            </GridItem>
            <GridItem area={"results"} pt={2}>
                <SimpleGrid minChildWidth='300px' spacing='30px' mt={'1%'} ml={'5%'} mr={'5%'}>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                    <FlightCard></FlightCard>
                </SimpleGrid>
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
                        View More Results
                    </Button>
                    {/*All results loaded*/}
                    <Text fontSize='4xl'>There are no more results to load</Text>
                </Box>
            </GridItem>
        </Grid>
    )
}

export default DiscoverFlights;