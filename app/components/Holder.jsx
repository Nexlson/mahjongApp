import Image from 'next/image'
import styled from 'styled-components';

const HolderTab = styled.div`
    margin: 5px 5px;
    @media (max-width: 500px) {
        width: 40px !important;
        height: 60px !important;
        margin: 5px 5px;
        padding: 0 !important;
    }
`

export default function Holder(props) {
    return(
        props.id == 0 ? (<HolderTab><Image src={"/grey.jpg"} alt={"Tile"+props.id} width={80} height={80}/></HolderTab>) : 
        (<HolderTab><Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.key} width={80} height={80}/></HolderTab>)
    )
}