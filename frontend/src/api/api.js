import sha256 from "crypto-js/sha256";

// TODO: change to read from env
const ENDPOINT = "http://localhost:8080";

const sendSignupReq = async (user, pass, email) => {
  const rawResponse = await fetch(ENDPOINT + "/signup", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: user,
      hashed_password: sha256(pass).toString(),
      email: email,
    }),
  });

  if (rawResponse.status !== 201) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

// returns a boolean indicating if a token is still valid for the session
const validateToken = async (tokenStr) => {
  if (tokenStr === "" || tokenStr === null) {
    return false;
  }

  const rawResponse = await fetch(ENDPOINT + "/validate-token", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      session_token: tokenStr,
    }),
  });
  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return false;
  }
  const content = await rawResponse.json();

  return content.valid;
};

const sendLoginReq = async (user, pass) => {
  const rawResponse = await fetch(ENDPOINT + "/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: user,
      hashed_password: sha256(pass).toString(),
    }),
  });

  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

const sendGenerateItineraryReq = async (
  session_token,
  startDateTime,
  endDateTime,
  cats
) => {
  const rawResponse = await fetch(ENDPOINT + "/generate-itinerary", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      session_token: session_token,
      start_time: Math.floor(startDateTime.getTime() / 1000),
      end_time: Math.floor(endDateTime.getTime() / 1000),
      preferred_categories: [...cats],
    }),
  });
  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

export { sendSignupReq, validateToken, sendLoginReq, sendGenerateItineraryReq };
