CREATE TABLE feeds.checkin (
    checkinid uuid PRIMARY KEY,
    category text,
    checkintimestamp timestamp,
    lat double,
    lng double,
    name text,
    placeid text,
    userid text
) 