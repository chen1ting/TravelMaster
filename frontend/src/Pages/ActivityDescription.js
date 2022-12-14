import {Box, Grid, GridItem, Heading, Stack} from '@chakra-ui/react';
import ActivityCard from '../Components/DiscoverComponents/ActivityCard';

// const DescriptionCard = () =>{
//     const Activity = {
//         title: 'Activity Title',
//         address:'Activity Address',
//         desc: 'Activity Description',

//     }
//     return (
//         <Stack maxW='sm' borderWidth='1px' borderRadius='lg' overflow='hidden' justify='center' align='center'>
//             <Box>
//                 {Activity.title}
//             </Box>
//             <Box>
//                 {Activity.address}
//             </Box>
//             <Box>
//                 {Activity.desc}
//             </Box>
//         </Stack>
//     )
// }

const ActivityDescription = () =>{
    return (
        <Grid
            h='600px'
            templateRows='repeat(2, 1fr)'
            templateColumns='repeat(7, 1fr)'
            gap={4}
            >
            <GridItem rowSpan={1} colSpan={2} bg='azure' justify="center" >
                <ActivityCard />
            </GridItem>
            <GridItem colSpan={5} bg='lightcyan' justify = 'center' alignContent='center'>
                Activity Description
            </GridItem>
            <GridItem colSpan={7} bg='azure'>
                View Activities Similar to this
            </GridItem>
        </Grid>
        
    )
}

export default ActivityDescription;