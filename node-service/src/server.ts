import express from "express";
import {players as data} from "./_mockdata_/data";
const app = express();
const port = 8080;

app.get("/players", (_req, res) => {

    res.send(data);
})

app.get("/", (_req, res) => {
    res.send({
        msg: "Visit /players to see mock data.",
        note: "Keep in mind this is test and for learning purposes."
    })
})

app.listen(port, () => {
    console.log("Server running on port: localhost:", port);
})