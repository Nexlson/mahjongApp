import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Grid from '@mui/material/Grid';
import styled from 'styled-components';
import NavButtons from './NavbarButton';
import { PrimaryColor } from '../utils/defaults'
import GitHubIcon from '@mui/icons-material/GitHub';
import Button from '@mui/material/Button';

const Title = styled(Grid)`
    width: 15vw;

    @media (max-width: 500px) {
        width: 10vw;
    }
`
const NavBut = styled(Button)`
    font-size: 1vw;
    margin-left: 5px;
    font-family: Roboto;
    maxHeight: 8vw;

    span {
        margin: 0;
    }

    @media (max-width: 500px) {
        font-size: 2vw;
        margin-left: 5px;
        font-family: Roboto;
        maxHeight: 6vw;
    }
`

export default function Navbar() {
    return (
        <AppBar position="static" style={{background: PrimaryColor, boxShadow: "none"}}>
            <Toolbar>
                <Title container>
                    <a href="/"><img src="/mahjong.ico" alt="Mahjong Logo" width={35} height={35}/></a>
                </Title>
                <Grid container direction="row" sx={{justifyContent: "flex-end", alignItems: "flex-end"}}>
                    <NavButtons name="Logger" link="/logger"/>
                    <NavButtons name="Calculator" link="/calculator"/>
                    <NavButtons name="Docs" link="/docs"/>
                    <NavBut variant="contained" startIcon={<GitHubIcon />} color="secondary" href="https://github.com/Nexlson"/>
                </Grid>
            </Toolbar>
        </AppBar>
    )
}