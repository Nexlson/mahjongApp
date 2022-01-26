import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Grid from '@mui/material/Grid';
import styled from 'styled-components';
import NavButtons from './NavbarButton';
import { PrimaryColor } from '../utils/defaults'

const Title = styled(Grid)`
    width: 15vw;

    @media (max-width: 500px) {
        width: 10vw;
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
                </Grid>
            </Toolbar>
        </AppBar>
    )
}