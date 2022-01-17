import Image from 'next/image'
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Toolbar from '@mui/material/Toolbar';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';

export default function Navbar() {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static" color="default">
                <Toolbar>
                    <Grid container sx={{flexGrow: 1}}>
                        <Image src="/mahjong.ico" alt="Mahjong Logo" width={30} height={20}/>
                        <Button color="inherit" href="/"><Typography variant="body1">Mahjong App</Typography></Button>
                    </Grid>
                    <Grid container sx={{justifyContent: "flex-end", alignItems: "flex-end"}}>
                        <Button color="inherit" variant="text" href="/"><Typography>Logger</Typography></Button>
                        <Button color="inherit" variant="text" href="/calculator"><Typography>Calculator</Typography></Button>
                        <Button color="inherit" variant="text" href="/docs"><Typography>Docs</Typography></Button>
                    </Grid>
                </Toolbar>
            </AppBar>
        </Box>
    )
}