import Image from 'next/image'
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Toolbar from '@mui/material/Toolbar';
import Grid from '@mui/material/Grid';

export default function Navbar() {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static" color="default">
                <Toolbar>
                    <Grid container sx={{flexGrow: 1}} className="title">
                        <a href="/"><Image src="/mahjong.ico" alt="Mahjong Logo" width={35} height={35}/></a>
                    </Grid>
                    <Grid container direction="row" sx={{justifyContent: "flex-end", alignItems: "flex-end"}}>
                        <Grid item>
                            <Button color="inherit" variant="contained" href="/" color={"secondary"} className="button"><strong>Logger</strong></Button>
                        </Grid>
                        <Grid item>
                            <Button color="inherit" variant="contained" href="/calculator" color={"secondary"} className="button"><strong>Calculator</strong></Button>
                        </Grid>
                        <Button color="inherit" variant="contained" href="/docs" color={"secondary"} className="button"><strong>Docs</strong></Button>
                    </Grid>
                </Toolbar>
            </AppBar>
        </Box>
    )
}