import {Button, Center, Text} from "@chakra-ui/react";

const LoadMoreRendering = (loadingMoreResults, allResultsShown) => {
    return (
        <Center m={10}>
            {/*Load more results button*/}
            {!allResultsShown &&
                <Button
                    isLoading={loadingMoreResults}
                    loadingText='Loading...'
                    colorScheme='teal'
                    variant='outline'
                >
                    {/*All results loaded*/}
                    View More Results
                </Button>
            }
            {/*No more results to load*/}
            {!loadingMoreResults && allResultsShown &&
                <Text fontSize='4xl'>There are no more results to load</Text>
            }
        </Center>
    )
}

export default LoadMoreRendering;