import sha256 from "crypto-js/sha256";

// TODO: change to read from env
const ENDPOINT = "http://localhost:8080";

const sendSignupReq = async (user, pass, email, pic) => {
  const formData = new FormData();
  formData.append("username", user);
  formData.append("hashed_password", sha256(pass).toString());
  formData.append("email", email);
  formData.append("avatar", pic);
  const rawResponse = await fetch(ENDPOINT + "/signup", {
    method: "POST",
    // headers: { "Content-Type": "multipart/form-data" },
    body: formData,
  });

  if (rawResponse.status !== 201) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

const sendCreateActivityReq = async (
  uid,
  title,
  isPaid,
  cats,
  desc,
  pic,
  hours
) => {
  const formData = new FormData();
  formData.append("user_id", uid);
  formData.append("title", title);
  formData.append("rating_score", 0);
  formData.append("paid", isPaid);
  for (var it = cats.values(), val = null; (val = it.next().value); ) {
    formData.append("category", val);
  }

  formData.append("description", desc);
  formData.append("longitude", 100); // TODO
  formData.append("latitude", 100);
  formData.append("image", pic);

  const days = ["sun", "mon", "tue", "wed", "thur", "fri", "sat"];
  for (let i = 0; i < days.length; i++) {
    formData.append(`${days[i]}_opening_time`, hours[i]);
    formData.append(`${days[i]}_closing_time`, hours[i] + 7);
  }

  const rawResponse = await fetch(ENDPOINT + "/create-activity", {
    method: "POST",
    // headers: { "Content-Type": "multipart/form-data" },
    body: formData,
  });

  if (rawResponse.status !== 201) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

const sendLogoutReq = async (session_token) => {
  const rawResponse = await fetch(ENDPOINT + "/logout", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      session_token: session_token,
    }),
  });

  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

const getActivityById = async (activityId, setActivity, setIsLoading) => {
  const rawResponse = await fetch(ENDPOINT + "/get-activity", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      activity_id: parseInt(activityId),
    }),
  });

  if (rawResponse.status !== 200) {
    setIsLoading(false);
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  setActivity(content);
  setIsLoading(false);
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
  return content.itinerary;
};

const getItinerary = async (
  id,
  session_token,
  setIsLoading,
  setNotifMsg,
  setTimeBins,
  setItineraryMap,
  setStartDate,
  setEndDate,
  setItineraryResp,
  setCurDate,
  setTimeBinsCopy,
  setItineraryMapCopy,
  setItiName
) => {
  const rawResponse = await fetch(ENDPOINT + "/get-itinerary", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id: id,
      session_token: session_token,
    }),
  });

  if (rawResponse.status === 401) {
    setNotifMsg("You do not have permission to view this itinerary!");
  } else if (rawResponse.status === 404) {
    setNotifMsg("Unfortunately, the page you are looking for does not exist.");
  }

  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }

  const content = await rawResponse.json();
  const itinerary = content.itinerary;
  if (itinerary.number_of_segments === 0) {
    setNotifMsg("Failed to generate an itinerary :(");
    return null;
  }
  setItineraryResp(itinerary);

  // TODO: some heavylifting work here
  // create as many 1 hour time bins as required
  // 1. Get the first epoch, determine its date (X).
  // 2. Get the last epoch, determine its date (Y).
  // 3. Then, start from 00:00 of X till 2400 of Y,
  // An expectation is as many time bins as required to easily build out the GUI.
  // Thus, we will additionally need to set date X and date Y in the state to facilitate the above.
  // Note: We should be able to do a one pass operation on the itinerary segments by calculating
  // the relative offset from the start date 00:00 to get the bucket index.
  // Also, we will need to maintain an auxilary hashmap to retrieve an activity summary by its id.

  // step 1:
  const firstDate = new Date(itinerary.start_time * 1000);
  const startDate = new Date(
    firstDate.getFullYear(),
    firstDate.getMonth(),
    firstDate.getDate(),
    0,
    0,
    0
  );
  setStartDate(startDate);
  setCurDate(startDate);

  // step 2:
  const lastDate = new Date(itinerary.end_time * 1000);
  const endDate = new Date( // end date is inclusive
    lastDate.getFullYear(),
    lastDate.getMonth(),
    lastDate.getDate(),
    0,
    0,
    0
  );
  setEndDate(endDate);

  // step 3:
  const buckets =
    Math.ceil((endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60)) +
    24; // 24 extra buckets for the last day which is inclusive

  // fill up the buckets
  const timeBins = [];
  for (let i = 0; i < buckets; i++) {
    timeBins.push(null);
  }

  // map each activity to its indexes in time bins
  const itineraryMap = new Map();

  for (let i = 0; i < itinerary.segments.length; i++) {
    const seg = itinerary.segments[i];
    let idx = (seg.start_time - startDate.getTime() / 1000) / (60 * 60);
    for (let cur = seg.start_time; cur < seg.end_time; cur += 60 * 60) {
      timeBins[idx] = seg.activity_summary;
      itineraryMap[seg.activity_summary.id] = (
        itineraryMap[seg.activity_summary.id] || []
      ).concat(idx);
      ++idx;
    }
  }
  console.log(timeBins);
  setTimeBins(timeBins);
  setItineraryMap(itineraryMap);
  setTimeBinsCopy([...timeBins]);
  setItineraryMapCopy(new Map(itineraryMap));
  setItiName(itinerary.name);
  setIsLoading(false);
  return content;
};

const getActivitiesByFilter = async (
  searchText,
  pageNum,
  pageSize,
  times,
  session_token,
  setActivities
) => {
  const rawResponse = await fetch(ENDPOINT + "/search-activity", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      search_text: searchText,
      page_no: pageNum,
      page_size: pageSize,
      times: times,
      session_token: session_token,
    }),
  });

  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }

  const content = await rawResponse.json();
  setActivities(content.activities);
};

const saveItineraryChanges = async (
  id,
  timeBins,
  session_token,
  startTime,
  itiName
) => {
  const segments = [];
  for (let i = 0; i < timeBins.length; i++) {
    if (timeBins[i] == null) {
      continue;
    }
    const act = timeBins[i];
    const t = startTime + i * 60 * 60;
    if (
      segments.length > 0 &&
      segments[segments.length - 1].activity_summary.id === act.id
    ) {
      segments[segments.length - 1].end_time = t + 60 * 60;
    } else {
      segments.push({
        start_time: t,
        end_time: t + 60 * 60,
        activity_summary: act,
      });
    }
  }
  const rawResponse = await fetch(ENDPOINT + "/save-itinerary", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id: id,
      segments: segments,
      session_token: session_token,
      name: itiName,
    }),
  });

  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }

  const content = await rawResponse.json();
  return content.id;
};

const getItisByUser = async (session_token, setItis) => {
  const rawResponse = await fetch(ENDPOINT + "/get-itineraries", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      session_token: session_token,
    }),
  });

  if (rawResponse.status !== 200) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }

  const content = await rawResponse.json();
  setItis(content.itineraries);
};

const addReview = async (
  session_token,
  aid,
  title,
  desc,
  stars,
  setActivity
) => {
  const rawResponse = await fetch(ENDPOINT + "/add-review", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      session_token: session_token,
      activity_id: aid,
      title: title,
      description: desc,
      rating: stars,
    }),
  });

  if (rawResponse.status !== 201) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return rawResponse.status;
  }

  const content = await rawResponse.json();
  setActivity(content);
  return 201;
};

export {
  sendSignupReq,
  validateToken,
  sendLoginReq,
  sendGenerateItineraryReq,
  getItinerary,
  getActivitiesByFilter,
  saveItineraryChanges,
  sendLogoutReq,
  getItisByUser,
  sendCreateActivityReq,
  getActivityById,
  addReview,
};
