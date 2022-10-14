import React, { useState } from "react";

import {
  Box,
  Button,
  Checkbox,
  Heading,
  Input,
  Link,
  Text,
  Textarea,
} from "@chakra-ui/react";
import { InfoIcon } from "@chakra-ui/icons";

import { submitFeedback } from "../api/api";

const Feedback = () => {
  const [feedbackMsg, setFeedbackMsg] = useState("");
  const [followUp, setFollowUp] = useState(false);

  const [afterSubmitMsg, setAfterSubmitMsg] = useState("");

  const sendSubmit = async () => {
    const data = await submitFeedback(
      window.sessionStorage.getItem("session_token"),
      feedbackMsg,
      followUp
    );
    if (data == null) {
      setAfterSubmitMsg(
        "Sorry, something went wrong creating your feedback. Please try again later."
      );
    } else {
      setAfterSubmitMsg(
        `Your feedback has been collected. ${
          followUp
            ? "As requested, we will follow up with your feedback within 3 business days"
            : ""
        }`
      );
    }
  };

  if (afterSubmitMsg) {
    return (
      <Box mt="16" textAlign="center">
        <Text>{afterSubmitMsg}</Text>
      </Box>
    );
  }

  return (
    <Box
      display="flex"
      flexDir="column"
      justifyContent="flex-start"
      alignItems="center"
      mt="10"
      rowGap="6"
    >
      <Heading>Help us improve!</Heading>

      <Box width="700px">
        <Text my="5">
          Our goal has always been to deliver the best possible experience in
          itinerary planning for your trip to Singapore.
          <br /> <br />
          If you have any feedback on how we could better serve you, we would
          love to hear from you!
        </Text>

        <Textarea
          h="200px"
          value={feedbackMsg}
          onChange={(e) => setFeedbackMsg(e.target.value)}
        />

        <Box my="8" display="flex" columnGap="5">
          <Text>
            Would you like us to follow up with you via email regarding your
            feedback?
            <Text fontSize="sm">
              We will get back to you within 3 business days.
            </Text>
          </Text>
          <Checkbox
            colorScheme="green"
            isChecked={followUp}
            onChange={(e) => setFollowUp(e.target.checked)}
          />
        </Box>

        <Box display="flex" mt="8">
          <InfoIcon mx="2" mt="0.5" />
          <Text fontSize="xs">
            By clicking on submit, you consent to us collecting any personal
            data you might have entered. We take your privacy seriously and will
            handle it with utmost due diligience. If you wish to delete any
            information you have previously submitted, kindly write to us at{" "}
            <Link color="teal.500" href="mailto:privacy@travelmaster.com">
              privacy@travelmaster.com
            </Link>
            .
          </Text>
        </Box>

        <Box display="flex" justifyContent="flex-end">
          <Button mt="2" mb="5" colorScheme="green" onClick={sendSubmit}>
            Submit
          </Button>
        </Box>
      </Box>
    </Box>
  );
};

export default Feedback;
