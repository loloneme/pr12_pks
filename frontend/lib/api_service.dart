import 'package:dio/dio.dart';
import 'package:frontend/models/cart_item.dart';
import './models/drink.dart';

class ApiService {
  final Dio _dio = Dio();

  final String _baseUrl = 'http://10.0.2.2:8080/handler';

  Future<List<Drink>> getDrinks([String? userID]) async {
    try {
      final url = userID != null ? '$_baseUrl/drink/all/$userID' : '$_baseUrl/drink/all';
      final response = await _dio.get(url);

      if (response.statusCode == 200) {
        List<Drink> drinks = (response.data as List)
            .map((drink) => Drink.fromJson(drink))
            .toList();

        drinks.sort((a, b) => a.id.compareTo(b.id));

        return drinks;
      } else {
        throw Exception('Failed to load drinks');
      }
    } catch (e) {
      throw Exception('Error fetching drinks: $e');
    }
  }

  Future<Drink> getDrinkByID(int drinkID) async {
    try {
      final response = await _dio.get('$_baseUrl/drink/$drinkID');
      if (response.statusCode == 200) {
        Drink drink = Drink.fromJson(response.data);

        drink.volumes?.sort((a, b) => a.price.compareTo(b.price));


        return drink;
      } else {
        throw Exception('Failed to load drink');
      }
    } catch (e) {
      throw Exception('Error fetching drink: $e');
    }
  }

  Future<void> updateDrink(Drink drink) async {
    try {
      final response = await _dio.put(
        '$_baseUrl/drink/${drink.id}',
          data:{
            'name': drink.name,
            'image': drink.image,
            'description': drink.description,
            'compound': drink.compound,
            'is_cold': drink.cold,
            'is_hot': drink.hot,
            'volumes': drink.volumes,
            'is_favorite': drink.isFavorite
          }
      );
      if (response.statusCode != 200) {
        throw Exception('Failed to update drink: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error updating drink: $e');
    }
  }


  Future<void> createDrink(Drink drink) async {
    try {
      final response = await _dio.post(
          '$_baseUrl/drink',
          data:{
              'name': drink.name,
            'image': drink.image,
            'description': drink.description,
            'compound': drink.compound,
            'is_cold': drink.cold,
            'is_hot': drink.hot,
            'volumes': drink.volumes,
            'is_favorite': drink.isFavorite
          }
      );
      if (response.statusCode == 201) {
        return;
      } else {
        throw Exception('Failed to create drink: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error creating drink: $e');
    }
  }

  Future<void> deleteDrink(int id) async {
    try {
      final response = await _dio.delete(
          '$_baseUrl/drink/$id',
      );
      if (response.statusCode != 200) {
        throw Exception('Failed to delete drink: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error deleting drink: $e');
    }
  }

  Future<void> toggleFavorite(int id, String userID) async {
    try {
      final response = await _dio.patch(
        '$_baseUrl/drink/$userID/$id',
      );
      if (response.statusCode != 200) {
        throw Exception('Failed to update drink status: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error updating drink status: $e');
    }
  }

  //Корзина
  Future<List<CartItem>> getCart(String userID) async {
    try {
      final response = await _dio.get(
        '$_baseUrl/cart/$userID',
      );
      if (response.statusCode == 200) {
        List<CartItem> cart = (response.data as List)
            .map((item) => CartItem.fromJson(item))
            .toList();

        cart.sort((a, b) => a.cartItemID.compareTo(b.cartItemID));

        return cart;
      } else {
        throw Exception('Failed to load cart');
      }
    } catch (e) {
      throw Exception('Error getting cart: $e');
    }
  }

  Future<void> updateCart(CartItem cartItem, String userID) async {
    try {
      final response = await _dio.patch(
          '$_baseUrl/cart/$userID/${cartItem.cartItemID}',
          data:{
            'quantity': cartItem.quantity,
          }
      );
      if (response.statusCode != 200) {
        throw Exception('Failed to update cart: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error updating cart: $e');
    }
  }

  Future<void> addToCart(CartItem cartItem, String userID) async {
    try {
      final response = await _dio.post(
          '$_baseUrl/cart/$userID',
          data:{
              "user_id": userID,
              "volume": {
                "volume_id": cartItem.volume.volumeID
              },
              "is_cold": cartItem.isCold,
              "is_hot": cartItem.isHot,
              "quantity": cartItem.quantity
          }
      );
      if (response.statusCode == 201) {
        return;
      } else {
        throw Exception('Failed to add to cart: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error adding to cart: $e');
    }
  }

  Future<void> deleteFromCart(int cartItemID, String userID) async {
    try {
      final response = await _dio.delete(
          '$_baseUrl/cart/$userID/$cartItemID'
      );
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to delete from cart: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Error deleting from cart: $e');
    }
  }
}