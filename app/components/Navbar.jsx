import Image from 'next/image'
import AppBar from '@mui/material/AppBar';
import Button from '@mui/material/Button';
import Toolbar from '@mui/material/Toolbar';
import Grid from '@mui/material/Grid';
import styled from 'styled-components';

const Title = styled(Grid)`
    width: 15vw;

    @media (max-width: 500px) {
        width: 10vw;
    }
`

const NavButton = styled(Button)`
    width: 8vw;
    font-size: 1vw;
    margin-left: 5px;
    font-family: Roboto;

    @media (max-width: 500px) {
        width: 2vw;
        font-size: 2vw;
        margin-left: 5px;
        font-family: Roboto;
      }
`

export default function Navbar() {
    return (
        <AppBar position="static" color="default">
            <Toolbar>
                <Title container sx={{flexGrow: 1}}>
                    <a href="/"><Image src="/mahjong.ico" alt="Mahjong Logo" width={35} height={35}/></a>
                </Title>
                <Grid container direction="row" sx={{justifyContent: "flex-end", alignItems: "flex-end"}}>
                    <Grid item>
                        <NavButton color="inherit" variant="contained" href="/" color={"secondary"}><strong>Logger</strong></NavButton>
                    </Grid>
                    <Grid item>
                        <NavButton color="inherit" variant="contained" href="/calculator" color={"secondary"}><strong>Calculator</strong></NavButton>
                    </Grid>
                    <NavButton color="inherit" variant="contained" href="/docs" color={"secondary"}><strong>Docs</strong></NavButton>
                </Grid>
            </Toolbar>
        </AppBar>
    )
}