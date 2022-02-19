import { Grid } from "@mui/material"
import { TilesBar } from "./TilesBar"

export default function Rule(props) {

    return (
        <Grid>
            <p> {props.name} </p>
            <p> {props.desc} </p>
            <TilesBar tiles={props.tiles}/>
        </Grid>
    )
}