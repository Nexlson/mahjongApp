import styled from 'styled-components';

const Image = styled.img`
    width: 60px;
    height: 80px;
    padding: 5px 5px 5px 5px;
    border: 5px hidden;

    @media (max-width: 500px) {
        width: 40px;
        height: 60px;
        margin: 2px 2px 2px 2px;
    }

`

const SelectedImage = styled(Image)`
    transform: rotate(90deg);
`

export default function Holder(props) {
    const imageClick = (idHolder) => {
        if (props.winningTile === idHolder){
            props.setWinningTile(0)
        }else{
            props.setWinningTile(idHolder)
        }
    }

    return(
        props.id === 0 ? (<Image src={"/grey.jpg"} alt={"Tile"+props.idHold}/>) : 
        (
            props.winningTile === props.idHold ? <SelectedImage src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.idHold} onClick={()=>imageClick(props.idHold)}/> :
            <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.idHold} onClick={()=>imageClick(props.idHold)}/>
        )
    )
}