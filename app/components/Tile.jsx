import Image from 'next/image'
import styled from 'styled-components';

const TileStyle = styled.div`
    margin: 5px 5px;
    :active {
        animation-name: selectTile;
        animation-duration: 1s;
    }
    @keyframes selectTile {
        from {width: 40px;}
        to {width: 50px;}
      }
      
    @media (max-width: 500px) {
        width: 40px;
        height: 40px;
    }
`

export default function Tile(props) {
    const imageClick = (id) => {
        props.setTiles([...props.tiles, id])
    }

    return(
        <TileStyle>
            <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.id} width={80} height={80} onClick={()=> imageClick(props.id)}/>
        </TileStyle>
    )
}