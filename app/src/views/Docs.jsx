import Navbar from "../components/Navbar"
import Rule from "../components/docs/Rule"
import { Grid } from "@mui/material"
import { rules } from "../utils/defaults"

export default function Docs(){
    return(
        <>
            <Navbar/>
            <Grid>
                {
                    rules.map((rule, index) => (
                        <Rule key={index} name={rule.name} desc={rule.desc} tiles={rule.tiles}/>
                    ))
                }
            </Grid>
        </>
    )
}