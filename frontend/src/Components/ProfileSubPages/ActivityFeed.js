import {Flex, Box,Grid,Stack, Button} from "@chakra-ui/react"

const ActivityFeed = () => {
    return (

    <Grid templateRows='repeat(6, 1fr)'
          gap={4}>
        <Flex justify="space-around">
            <Button>
                Post a picture
            </Button>
            <Button>
                Write a review
            </Button>
        </Flex>
        <Box>
            Review1
        </Box>
        <Box>
            Review2
        </Box>
        <Box>
            Review3
        </Box>

    </Grid>


)
}


export default ActivityFeed;