import { Button, DialogActions, DialogContent, DialogTitle, TextField, Select, MenuItem, InputLabel, FormControl, ToggleButtonGroup, ToggleButton } from "@mui/material";
import { useState } from "react"


function CreateRoomContent(){
    const [count, setCount] = useState('');
    const [mode, setMode] = useState('');

    const handleCountChange = (event) => {
        setCount(event.target.value);
    }

    const handleModeChange = (event) => {
        setMode(event.target.value);
    }

    return(
        <>
            <TextField autoFocus id="Room Name" label="Room Name" />
            <TextField autoFocus id="Passcode" label="Passcode" />
            <FormControl sx={{ minWidth: 120 }}>
                <InputLabel id="select-label">Players</InputLabel>
                <Select value={count} labelId="select-label" id="Player Count" label="Player Count" onChange={handleCountChange}>
                    <MenuItem value={3}> 3 </MenuItem>
                    <MenuItem value={4}> 4 </MenuItem>
                </Select>
            </FormControl>
            <ToggleButtonGroup color="primary" value={mode} onChange={handleModeChange}>
                <ToggleButton value="Normal">Normal</ToggleButton>
                <ToggleButton value="Interactive">Interactive</ToggleButton>
            </ToggleButtonGroup>
        </>
    )
}

function JoinRoomContent(){
    const [role, setRole] = useState('');

    const handleRoleChange = (event) => {
        setRole(event.target.value);
    }

    return(
        <>
            <TextField autoFocus id="Room Name" label="Room Name" />
            <TextField autoFocus id="Passcode" label="Passcode" />

            <InputLabel id="select-label">Role</InputLabel>
            <ToggleButtonGroup color="primary" value={role} onChange={handleRoleChange}>
                <ToggleButton value="King">King</ToggleButton>
                <ToggleButton value="Queen">Queen</ToggleButton>
                <ToggleButton value="Jack">Jack</ToggleButton>
            </ToggleButtonGroup>

        </>
    )
}

export default function PopUpContent(props) {

    return(
        <>
            <DialogTitle> {props.name} a game! </DialogTitle>
            <DialogContent>
                {
                    props.create ?
                    <CreateRoomContent />:
                    <JoinRoomContent/>
                }
            </DialogContent>
            <DialogActions>
                <Button onClick={props.closeDialog}>{props.name}</Button>
                <Button onClick={props.closeDialog}>Cancel</Button>
            </DialogActions>
        </>
    )
}