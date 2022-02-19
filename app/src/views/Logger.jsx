import GameButton from "../components/logger/GameButton"
import Navbar from "../components/Navbar"
import PopUpContent from "../components/logger/PopUpContent"
import { useState } from "react"
import { Button, Dialog, DialogActions, DialogContent, DialogTitle } from "@mui/material";

export default function Logger(){
    const [popUp, setpopUp] = useState(false);
    const [create, setCreate] = useState(false);
    const [join, setJoin] = useState(false);

    const closeDialog = () => {
        setpopUp(false);
        setCreate(false);
        setJoin(false);
    }

    return(
        <>
            <Navbar/>
            <GameButton desc="Create" upper={true} click={setpopUp} state={setCreate}/>
            <Dialog open={popUp} onClose={closeDialog}>
                {
                    create ?
                    <PopUpContent name="Create" closeDialog={closeDialog} create={create}/>
                    :
                        ""
                }
                {
                    join? 
                    <PopUpContent name="Join" closeDialog={closeDialog}/>
                    :
                        ""
                }

            </Dialog>
            <GameButton desc="Join" click={setpopUp} state={setJoin}/>
        </>
    )
}