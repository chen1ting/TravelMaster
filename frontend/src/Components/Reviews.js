import {useState} from 'react';
//import { useAuth } from '../lib/auth';
import {
    FormControl, FormLabel, Button, Input, Box, Flex, Grid, GridItem
} from '@chakra-ui/react';
import {
    MDBCard,
    MDBCardImage,
    MDBCardBody,
    MDBCardTitle,
    MDBCardText,
    MDBRow,
    MDBCol
} from 'mdb-react-ui-kit';
const fields_width = '52.5%';
let text;
text = "Penguin Feeding Show"
const reviewList = [
    { name: "John", age: 24 },
    { name: "Linda", age: 19 },
    { name: "Josh", age: 33 }
];
const reviewCount = 3;
const renderReviews = () => {
    const result = [];
    for (let i = 0; i < reviewCount; i++) {
        <MDBCard>
            <MDBCardBody>Hello</MDBCardBody>
        </MDBCard>
    }

    return <ul>{result}</ul>;
};

const Reviews = () => {

    //const { signIn } = useAuth();


    return (<Grid
        templateAreas={`"left_top right"
                            "left_bottom right"
                            `}
        gridTemplateRows={'50fr 50fr'}
        gridTemplateColumns={'1fr 2fr'}
        h='650px'
        // gap='1'
        // color='blackAlpha.700'
        fontWeight='bold'
    >
        <GridItem pl='2' bg='blue.200' area={'left_top'}>
            <h1>
                <font size={6}>
                    Reviews
                </font>
            </h1>

            <MDBCard>
                <MDBCardImage src='https://mdbootstrap.com/img/new/standard/nature/182.webp' alt='...' position='top' />
                <MDBCardBody>
                    <MDBCardText>
                        {text}
                    </MDBCardText>
                </MDBCardBody>
            </MDBCard>
        </GridItem>
        <GridItem pl='2' bg='white' area={'left_bottom'}>

        </GridItem>
        <GridItem pl='2' bg='white' area={'right'}>
            <MDBRow className='row-cols-1 row-cols-md-3 g-4'>
                <MDBCol>
                    <MDBCard>
                        <MDBCardImage
                            src='https://mdbootstrap.com/img/new/standard/city/041.webp'
                            width={250} height={250}
                            alt='...'
                            position='top'
                        />
                        <MDBCardBody>
                            <MDBCardTitle>Card title</MDBCardTitle>
                            <MDBCardText>
                                This is a longer card with supporting text below as a natural lead-in to additional content.
                                This content is a little bit longer.
                            </MDBCardText>
                        </MDBCardBody>
                    </MDBCard>
                </MDBCol>
                <MDBCol>
                    <MDBCard>
                        <MDBCardImage
                            src='https://mdbootstrap.com/img/new/standard/city/042.webp'
                            width={250} height={250}
                            alt='...'
                            position='top'
                        />
                        <MDBCardBody>
                            <MDBCardTitle>Card title</MDBCardTitle>
                            <MDBCardText>
                                This is a longer card with supporting text below as a natural lead-in to additional content.
                                This content is a little bit longer.
                            </MDBCardText>
                        </MDBCardBody>
                    </MDBCard>
                </MDBCol>
                <MDBCol>
                    <MDBCard>
                        <MDBCardImage
                            src='https://mdbootstrap.com/img/new/standard/city/043.webp'
                            width={250} height={250}
                            alt='...'
                            position='top'
                        />
                        <MDBCardBody>
                            <MDBCardTitle>Card title</MDBCardTitle>
                            <MDBCardText>
                                This is a longer card with supporting text below as a natural lead-in to additional content.
                                This content is a little bit longer.
                            </MDBCardText>
                        </MDBCardBody>
                    </MDBCard>
                </MDBCol>
                <MDBCol>
                    <MDBCard>
                        <MDBCardImage
                            src='https://mdbootstrap.com/img/new/standard/city/044.webp'
                            width={250} height={250}
                            alt='...'
                            position='top'
                        />
                        <MDBCardBody>
                            <MDBCardTitle>Card title</MDBCardTitle>
                            <MDBCardText>
                                This is a longer card with supporting text below as a natural lead-in to additional content.
                                This content is a little bit longer.
                            </MDBCardText>
                        </MDBCardBody>
                    </MDBCard>
                </MDBCol>
            </MDBRow>

        </GridItem>
    </Grid>);
};

export default Reviews;