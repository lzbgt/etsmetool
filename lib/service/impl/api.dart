import 'dart:async';
//import 'package:http/http.dart' as http;
import 'package:dio/dio.dart';
import 'package:etstool/model/home/func.dart';
import 'dart:convert';
import '../api.dart';

import 'package:hive/hive.dart';
import 'package:injector/injector.dart';
import '../../const/consts.dart';
import '../../model/common/message.dart';
import 'package:flutter/material.dart';

class ProdAPI implements API {
  const ProdAPI({this.baseUrl = 'http://121.41.121.222:9090'});
  final String baseUrl;

  @override
  Future<RespMessage<String>> login(String phone, String password) async {
    final cst = Consts();
    final box = Injector.appInstance.get<Box>();
    var tok = box.get(cst.authTokenKey);
    print('previous tok: $tok');
    tok = '';
    try {
      box.put(cst.authTokenKey, '');
      final resp =
          await Dio(BaseOptions(connectTimeout: 3000, receiveTimeout: 3000))
              .get(baseUrl + '/login/$phone/$password');
      print(resp);
      tok = resp.data['token'];
      print('new tok: $tok');
      box.put(cst.authTokenKey, tok.toString());
    } catch (e) {
      print('exception login: $e');
      return RespMessage(code: 1, message: e.toString());
    } finally {}

    return RespMessage(code: 0, data: tok);
  }

  @override
  List<FuncItemData> getFuncMods() {
    return <FuncItemData>[
      FuncItemData(title: 'Query Info', icon: Icons.info, index: 0),
      FuncItemData(title: 'Unbind Device', icon: Icons.unsubscribe, index: 1),
    ];
  }
}
