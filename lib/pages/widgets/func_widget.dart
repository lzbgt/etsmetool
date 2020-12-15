import 'package:etstool/bloc/home/home_bloc.dart';
import 'package:etstool/model/common/api_dt.dart';
import 'package:flutter/material.dart';
import '../../model/home/func.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

typedef dynamic FuncItemDataCallBack(FuncItemData value);

abstract class FuncWidget extends StatelessWidget {
  FuncWidget({
    Key key,
    @required this.index,
    @required this.itemData,
    @required this.isShort,
    @required this.bloc,
    // this.onTap,
    // this.onSubmmit,
  }) : super(key: key);
  final int index;
  final FuncItemData itemData;
  final bool isShort;
  final _textController = TextEditingController();
  final Bloc bloc;
  // final FuncItemDataCallBack onTap;
  // final FuncItemDataCallBack onSubmmit;

  /// has to be provided by concrete class
  dynamic onTap(FuncItemData value);

  dynamic onSubmit(FuncItemData value, String formValue);

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 6.0,
      color: Colors.lightGreen.shade50,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
      ),
      margin: EdgeInsets.all(8.0),
      child: InkWell(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ListTile(
              leading: Icon(
                itemData.icon,
                size: 48.0,
                color: Colors.lightGreen,
              ),
              title: Text(itemData.title),
              subtitle: TextFormField(
                controller: _textController,
              ),
              trailing: RaisedButton(
                onPressed: () {
                  print('submit pressed');
                  onSubmit(itemData, _textController.text);
                },
                child: Text('Submit'),
              ),
            ),
            Divider(),
            Text(
              itemData.title,
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: 16.0,
              ),
            )
          ],
        ),
        onTap: () {
          print('idx $index');
          // Navigator.pop(context, itemData.index);
          onTap(itemData);
        },
      ),
    );
  }
}

class UserInfoFuncWidget extends FuncWidget {
  UserInfoFuncWidget({
    int index,
    Bloc bloc,
    FuncItemData itemData,
    bool isShort,
  }) : super(
          index: index,
          bloc: bloc,
          itemData: itemData,
          isShort: isShort,
        );

  @override
  onSubmit(FuncItemData item, String formValue) {
    bloc.add(FuncModSubmitted(
      index: index,
      req: UserDeviceInfoRequest(env: 'prod', field: formValue),
    ));
  }

  @override
  onTap(FuncItemData value) {}
}
