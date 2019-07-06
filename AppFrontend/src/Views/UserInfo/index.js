import React, { Component } from 'react';
import { Text, View, FlatList } from 'react-native';
import { ListItem, Avatar } from "react-native-elements";

export default class UserInfoScreen extends Component {

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>个人信息</Text>)
    };

    render() {
        return (
            <View>
                <View style={{height: 200}}>
                    <View style={{alignItems: 'center', justifyContent: 'center'}}>
                        <View style={{height: 55}}/>
                        <Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')} />
                        <Text style={{fontSize: 15}}>点击修改头像</Text>
                    </View>
                </View>
                <ListItem
                    topDivider
                    bottomDivider
                    chevron
                    title={'用户名'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>chiangel</Text>}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'真实姓名'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>林江浩</Text>}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'性别'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>男</Text>}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'邮箱'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>chiangel.ljh@gmail.com</Text>}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'手机'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>15221278083</Text>}
                />
            </View>
        )
    }
}