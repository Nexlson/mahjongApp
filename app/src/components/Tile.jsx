import styled from 'styled-components';

const TileStyle = styled.div`
    margin: 3px 3px;
    @media (max-width: 500px) {
        width: 30px;
        height: 30px;
    }
`

const Image = styled.img `
    width: 60px;
    height: 80px;

    @media (max-width: 500px) {
        width: 30px;
        height: 50px;
    }

    :active {
        animation-name: selectTile;
        animation-duration: 0.1s;
    }

    @keyframes selectTile {
        from {transform:translateY(0);}
        to {transform:translateY(-10px);}
    }
`

export default function Tile(props) {
    const imageClick = (id) => {
        props.setTiles([...props.tiles, id])
    }

    return(
        <TileStyle>
            <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.id} onClick={()=> imageClick(props.id)}/>
        </TileStyle>
    )
}