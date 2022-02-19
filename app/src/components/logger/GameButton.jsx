import { Button, Grid } from "@mui/material";
import styled from 'styled-components';

const BigButton = styled(Button)`
    @media (max-width: 500px) {
        width: 50vw;
        height: 50vw;

        :active{
            background-color: red;
        }
      }

`

const UpperGrid = styled(Grid)`
    @media (max-width: 500px) {
        padding-top: 20vw; 
    }
`

const LowerGrid = styled(Grid)`
    @media (max-width: 500px) {
        padding-top: 20vw; 
        float: right;
    }
`
function CustomButton(props){
    return(
        <BigButton variant="contained" onClick={()=> {props.click(true); props.state(true)}}>
            {props.desc}
        </BigButton>
    )
}

export default function GameButton(props) {
    return(
        <>
            {
            props.upper ?          
            <UpperGrid>
                <CustomButton click={props.click} desc={props.desc} state={props.state}/>
            </UpperGrid> :
            <LowerGrid>
                <CustomButton click={props.click} desc={props.desc} state={props.state}/>
            </LowerGrid>
            }
        </>
    )
}