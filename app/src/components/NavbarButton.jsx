import Grid from '@mui/material/Grid';
import styled from 'styled-components';
import Button from '@mui/material/Button';

const NavButton = styled(Button)`
    width: 8vw;
    font-size: 1vw;
    margin-left: 5px;
    font-family: Roboto;
    heigth: 8vw;

    @media (max-width: 500px) {
        width: 2vw;
        font-size: 2vw;
        margin-left: 5px;
        font-family: Roboto;
      }
`

export default function NavButtons(props) {

    return(
        <Grid item>
            <NavButton color="secondary" variant="contained" href={props.link} >
                <strong>{props.name}</strong>
            </NavButton>
        </Grid>
    )
}